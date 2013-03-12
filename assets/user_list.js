$(document).ready(function() {
    $("a[a_remove_link='true']").bind('click', function() {
        var link = $(this);
        $.ajax('/profile?email=' + link.attr('email-val'), {
            type: 'DELETE'
        }).done(function() {
            link.parent().parent().remove();
        });
    });
});