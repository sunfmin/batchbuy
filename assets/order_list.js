$(document).ready(function() {
    clip = new ZeroClipboard($("#copy-button"), { 
        moviePath: "/assets/ZeroClipboard.swf",
        hoverClass: "bootstrap-btn-hover",
        activeClass: "bootstrap-btn-active"
    });
    // $('#copy-button-view').bind('click', function() {
    // $('#copy-button').click();
    // console.log('clicking');
    // })
});
