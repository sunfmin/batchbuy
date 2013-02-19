$(document).ready(function() {
    $('input[type=email]').val($.cookie('email'));
    $('input[type=date]').val(moment(new Date).format('YYYY-MM-DD'));
    
    // disabled the default enter-key-pressing event on all the forms
    $('form').keypress( function(event){ if (event.which == '13') { event.preventDefault(); } });
    
    if (moment($('#current-date').html()).isBefore(moment().format('YYYY-MM-DD'))) { 
        $('button').attr('disabled', true);
        $('input').attr('disabled', true);
        return;
    }
    
    var getOrder = function(inputField) {
        return {
            email: $('#email').val(),
            date: $('#current-date').html(),
            productid: inputField.parent().parent().find('input[type="hidden"]').val(),
            count: inputField.val()
        };
    };
    
    var putOrder = function(inputField, doneCallback) {
        $.ajax('/order', {
            type: 'POST',
            data: getOrder(inputField)
        }).done(doneCallback);
    };
    
    // on count field changing
    $('input[type=number]').on('change', function() {
        var input = $(this),
            val = parseInt(input.val());
        
        // var field = $(this);
        input.attr('data-original-title', 'done');
        input.tooltip('show');
        input.attr('data-original-title', '');
        var doneCallback = function() {
            setTimeout(function() { input.tooltip('hide'); }, 1000);
        };

        if (val <= 0) { 
            input.val(0);
            $.ajax('/order?' + $.param(getOrder(input)), { type: 'DELETE' }).done(doneCallback);
        } else {
            putOrder(input, doneCallback);
        }
    });
    
    // on clicking plus button
    $('.input-prepend .btn[name="plus"]').on('click', function() {
        var input = $(this).parent().find('input'); 
        input.val(parseInt(input.val()) + 1);
        
        var button = $(this);
        button.attr('data-original-title', 'done');
        button.tooltip('show');
        button.attr('data-original-title', '');
        var doneCallback = function() {
            setTimeout(function() { button.tooltip('hide'); }, 1000);
        };
        
        putOrder(input);
    });
    
    // on clicking minus button
    $('.input-prepend .btn[name="minus"]').on('click', function() {
        var input = $(this).parent().find('input'),
            val = parseInt(input.val()); 
        
        var button = $(this);
        button.attr('data-original-title', 'done');
        button.tooltip('show');
        button.attr('data-original-title', '');
        var doneCallback = function() {
            setTimeout(function() { button.tooltip('hide'); }, 1000);
        };
            
        if (val && val - 1) { 
            input.val(val - 1);
            putOrder(input, doneCallback);
        } else {
            input.val(0);
            $.ajax('/order?' + $.param(getOrder(input)), { type: 'DELETE' }).done(doneCallback);
        }
    });
});