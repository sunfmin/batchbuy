$(document).ready(function() {
    var currentUser = $.cookie('email'),
        managable = ['venustingting@gmail.com', 'bom.d.van@gmail.com'].indexOf(currentUser) != -1;

    if (!managable) {
        $("#copy-button").remove();
        return;
    };

    var clip = new ZeroClipboard($("#copy-button"), {
        moviePath: "/assets/ZeroClipboard.swf",
        hoverClass: "bootstrap-btn-hover",
        activeClass: "bootstrap-btn-active"
    });

    clip.on('complete', function() {
    	$.ajax('/no_more_order_today', {
    		method: 'POST',
    		data: { date: $('#current-date').text() }
    	}).done(function() {
            $('#allow-new-order-again').show();
        });
    });

    // no more order today
     $.ajax({
        url: 'is_no_more_order_today?date=' + $('#current-date').text()
    }).done(function(response) {
        if (response != 'true') { return; };

        $('#allow-new-order-again').show();
    });

    $('#allow-new-order-again').click(function() {
        $.ajax('/make_more_order_today', {
            method: 'POST',
            data: { date: $('#current-date').text() }
        }).done(function() {
            $('#allow-new-order-again').hide();
        });
    });
});
