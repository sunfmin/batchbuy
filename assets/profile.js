$(document).ready(function() {
    $("input[name=email]").val($.cookie("email"));
    $("input[name=name]").val($.cookie("name"));
    $("input[name=avatarlink]").val($.cookie("avatarlink"));
    
	$("form").submit(function() {
		var submitable = true;
		submitable = !!$("input[name=name]").val();
		submitable = !!$("input[name=email]").val();
		
		if (submitable) {
		    $.cookie("email", $("input[name=email]").val());
		    $.cookie("name", $("input[name=name]").val());
		    $.cookie("avatarlink", $("input[name=avatarlink]").val());
		}
		else alert("Name and Email Can't be EMPTY!");
		return submitable;
	})
});

// var UserModel = Backbone.Model.extend({
// 	urlRoot: '/profile.json'
// });
// 
// var user = new UserModel();
// 
// user.save({
// 	Name: "user",
// 	Email: "user@test.com"
// }, {
// 	success: function(user) {
// 		console.log(user.toJson());
// 	}
// });