$(document).ready(function() {
    $('input[type=email]').val($.cookie('email'));
    $('input[type=date]').val(moment(new Date).format('YYYY-MM-DD'));
    
    // disabled the default enter-key-pressing event on all the forms
    $('form').keypress( function(event){ if (event.which == '13') { event.preventDefault(); } });
    
    var getOrder = function(inputField) {
        return {
            email: $('#email').val(),
            date: $('#current-date').html(),
            productid: inputField.parent().parent().find('input[type="hidden"]').val(),
            count: inputField.val()
        }
    };
    
    var putOrder = function(inputField) {
        $.ajax('/order', {
            type: 'POST',
            data: getOrder(inputField)
        });
    };
    
    // on count field changing
    $('input[type=number]').on('change', function() {
        var input = $(this),
            val = parseInt(input.val());
        
        if (val <= 0) { 
            input.val(0);
            $.ajax('/order?' + $.param(getOrder(input)), { type: 'DELETE' });
        } else {
            putOrder(input);
        }
    });
    
    // on clicking plus button
    $('.input-prepend .btn[name="plus"]').on('click', function() {
        var input = $(this).parent().find('input'); 
        input.val(parseInt(input.val()) + 1);
        
        putOrder(input);
    });
    
    // on clicking minus button
    $('.input-prepend .btn[name="minus"]').on('click', function() {
        var input = $(this).parent().find('input'),
            val = parseInt(input.val()); 
        
        if (val && val - 1) { 
            input.val(val - 1);
            putOrder(input);
        } else {
            input.val(0);
            $.ajax('/order?' + $.param(getOrder(input)), { type: 'DELETE' });
        }
    });
    
    // put order
    // $('#submitBtn').click(function() {
    //     var reloadCount = 0;
    //     if ($('input[type=email]').val() == '') {
    //         alert('Please enter your email before submit your orders.');
    //         return;
    //     };
    //     $('input[type="number"]').each(function(index, count) {
    //         if ($(count).val() != "") {
    //             reloadCount += 1;
    //             var data = { email: $('input[type=email]').val()};
    //             $(count).parent().parent().find('input').each(function(index, input) {
    //                 data[$(input).attr('name')] = $(input).val();
    //             });
    //             
    //             
    //             $.ajax('/order', {
    //                 type: 'POST',
    //                 data: data
    //             }).done(function() {
    //                 reloadCount -= 1;
    //                 if (reloadCount == 0) {location.reload();};
    //             });
    //         };
    //     });
    // });
    
});