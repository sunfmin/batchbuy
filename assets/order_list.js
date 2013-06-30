$(document).ready(function() {
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

    var today = moment(new Date()),
        currentPageDate = moment($('#current-date').text(), 'YYYY-MM-DD'),
        sameDate = today.format('YYYY-MM-DD') == currentPageDate.format('YYYY-MM-DD');

    if (document.isNoMoreOrderToday && sameDate) {
        $('#allow-new-order-again').show();
    }

    $('#allow-new-order-again').click(function() {
        $.ajax('/make_more_order_today', {
            method: 'POST',
            data: { date: $('#current-date').text() }
        }).done(function() {
            $('#allow-new-order-again').hide();
        });
    });
});
