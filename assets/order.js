$(document).ready(function() {
    $('input[type=email]').val($.cookie('email'));
    $('input[type=date]').val(moment(new Date).format('YYYY-MM-DD'));
    
    $('#submitBtn').click(function() {
        var reloadCount = 0;
        if ($('input[type=email]').val() == '') {
            alert('Please enter your email before submit your orders.');
            return;
        };
        $('input[type="number"]').each(function(index, count) {
            if ($(count).val() != "") {
                reloadCount += 1;
                var data = { email: $('input[type=email]').val()};
                $(count).parent().parent().find('input').each(function(index, input) {
                    data[$(input).attr('name')] = $(input).val();
                });
                
                
                $.ajax('/order', {
                    type: 'POST',
                    data: data
                }).done(function() {
                    reloadCount -= 1;
                    if (reloadCount == 0) {location.reload();};
                });
            };
        });
    });
    
});