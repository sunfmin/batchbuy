$(document).ready(function() {
	$('input[name=email]').val($.cookie('email'));
	
	$('img').attr('src', getCurrentProduct().PhotoLink);
	$('select').bind('change', function() {
		$('img').attr('src', getCurrentProduct().PhotoLink);
		console.log(getCurrentProduct().PhotoLink);
	})
});

function getCurrentProduct () {
	var currentProductId = $('select').val();
	
	return products.filter(function(product) {
		return product.Id == currentProductId
	})[0];
}