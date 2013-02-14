$(document).ready(function() {
    var a = $('a[href="'+ window.location.pathname+ '"]');
    a.parent().addClass('active');
    document.title = $('a[href="'+ window.location.pathname+ '"]').html();
});