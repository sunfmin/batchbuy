$(document).ready(function() {
	$("select").bind('change', function() {
		var productId = $(this).val(),
			product = products.filter(function(product) {
				return product.Id == productId;
			})[0];
		if (!product) { product = {ValidFrom: new Date(), ValidTo: new Date()}; }
		$("input[name=name]").val(product.Name);
		$("input[name=photolink]").val(product.PhotoLink);
		$("input[name=price]").val(product.Price);
		$("input[name=validfrom]").val(moment(product.ValidFrom).format('YYYY-MM-DD'));
		$("input[name=validto]").val(moment(product.ValidTo).format('YYYY-MM-DD'));
	});
});