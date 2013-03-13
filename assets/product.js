$(document).ready(function() {
    $('#new-product-btn').bind('click', function() {
        $("html, body").animate({ scrollTop: $('html, body').height() }, "slow")
    });

    $('#back-to-page-top-btn').bind('click', function() {
        $("html, body").animate({ scrollTop: 0 }, "slow")
    });

    $('div[product-state]').each(function() {
        var validFrom = moment($(this).parent().find('input[name=validfrom]').val()),
            validTo = moment($(this).parent().find('input[name=validto]').val());

        if (validFrom.isSame('0001-01-01', 'year') && validTo.isSame('0001-01-01', 'year')) {
            $(this).html('有货');
        } else {
            $(this).html('缺货');
            $(this).addClass('btn-warning');
            $(this).attr('product-state', 'false');
        }
    });

    var postUpdate = function() {
        var data = {},
            captionPanel = $(this).parent().parent();
        captionPanel.find('input').each(function() {
            data[$(this).attr('name')] = $(this).val();
        });

        $.ajax('/product', {
            type: 'POST',
            data: data
        });
    };

    var updateImg = function() {
        var img = $(this).parent().parent().parent().find('img');
        img.attr('src', $(this).val());
        console.log("new PhotoLink: " + $(this).val());
    };

    var updateProductState = function() {
        if ($(this).attr('product-state') == "true") {
            $(this).parent().find('input[name=validfrom]').val('1970-01-01');
            $(this).parent().find('input[name=validto]').val('1970-01-01');
            $(this).attr('product-state', 'false');
            $(this).html('缺货');
            $(this).addClass('btn-warning');
        } else {
            $(this).parent().find('input[name=validfrom]').val('0001-01-01');
            $(this).parent().find('input[name=validto]').val('0001-01-01');
            $(this).attr('product-state', 'true');
            $(this).html('有货');
            $(this).removeClass('btn-warning');
        }

        $(this).parent().find('input[name=validto]').first().change();
    };

    var insertNewProductForm = function () {
        var templ = "<li class='span4'><form class='thumbnail'>\
            <img src=''>\
            <div class='caption'>\
                <div class='input-prepend'>\
                  <span class='add-on product-add-on'>Name</span>\
                  <input class='span2 product-input' type='text' name='name' value=''>\
                </div>\
                <div class='input-prepend'>\
                  <span class='add-on product-add-on'>Price</span>\
                  <input class='span2 product-input' type='number' name='price' value=''>\
                </div>\
                <div class='input-prepend'>\
                  <span class='add-on product-add-on'>PhotoLink</span>\
                  <input class='span2 product-input' type='text' name='photolink' value=''>\
                </div>\
                <input type='hidden' name='validfrom' value='0001-01-01'>\
                <input type='hidden' name='validto' value='0001-01-01'>\
                <div class='btn btn-primary pull-right product-state-btn' product-state='true'>有货</div>\
                <div style='clear:both;'></div>\
                <div id='savingDevice'>\
                    <a class='btn btn-primary pull-right'>Save</a>\
                    <div style='clear:both;'></div>\
                </div>\
            </div>\
        </form></li>";

        var lastRow = $('#main-panel').find('ul').last(),
            lastRowAppendable = lastRow.find('li').length < 3;
        if (!lastRowAppendable) {
            lastRow.after("<div class='row-fluid'><ul class='thumbnails'></ul></div>");
            lastRow = $('#main-panel').find('ul').last();
        };

        lastRow.append(templ);
        var currentForm = lastRow.find('form').last();
        currentForm.find('input[name="photolink"]').bind('change', updateImg);
        $('div[product-state]').last().bind('click', updateProductState);
        // $(".datepicker").datepicker({ dateFormat: 'yy-mm-dd' });

        lastRow.find('li').last().find('a').bind('click', function() {
            var data = {},
                captionPanel = $(this).parent().parent();
            captionPanel.find('input').each(function() {
                data[$(this).attr('name')] = $(this).val();
            });

            $.ajax('/product', {
                type: 'POST',
                data: data
            }).done(function(product) {
                captionPanel.prepend("<input type='hidden' name='productid' value='" + $.parseJSON(product).Id + "'>");
                $("#savingDevice").remove();
                currentForm.find('input').bind('change', postUpdate);

                insertNewProductForm();
            });
        });
    };

    // $(".datepicker").datepicker({ dateFormat: 'yy-mm-dd' });
    $('div[product-state]').bind('click', updateProductState);

    $("input").bind('change', postUpdate);
    $('input[name="photolink"]').bind('change', updateImg);

    insertNewProductForm();
});