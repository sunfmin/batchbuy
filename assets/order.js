$(document).ready(function() {
    $('input[type=email]').val($.cookie('email'));
    $('input[type=date]').val(moment(new Date).format('YYYY-MM-DD'));
    
    // monitor plus button
    $('.input-prepend .btn[name="plus"]').on('click', function() {
        var input = $(this).parent().find('input'); 
        input.val(parseInt(input.val()) + 1);
        
        $.ajax('/order', {
            type: 'PUT',
            data: {
                email: $('#email').val(),
                date: $('#current-date').html(),
                productid: $(this).parent().parent().find('input[type="hidden"]').val(),
                count: input.val()
            }
        });
    });
    
    // monitor plus button
    $('.input-prepend .btn[name="minus"]').on('click', function() {
        var input = $(this).parent().find('input'),
            val = parseInt(input.val()); 
        
        if (val) { 
            input.val(val - 1);
            
            $.ajax('/order', {
                type: 'PUT',
                data: {
                    email: $('#email').val(),
                    date: $('#current-date').html(),
                    productid: $(this).parent().parent().find('input[type="hidden"]').val(),
                    count: input.val()
                }
            });
        } else {
            $.ajax('/order?' + $.param({
                email: $('#email').val(),
                date: $('#current-date').html(),
                productid: $(this).parent().parent().find('input[type="hidden"]').val()
            }), {
                type: 'DELETE'
            });
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