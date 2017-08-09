//////CONTACT FORM VALIDATION
jQuery(document).ready(function ($) {
	
	//if submit button is clicked
	$('#submit').click(function () {		
		
		//Get the data from all the fields
		var name = $('input[name=name]');
		var email = $('input[name=email]');
		var regx = /^([a-z0-9_\-\.])+\@([a-z0-9_\-\.])+\.([a-z]{2,4})$/i;
		var returnError = false;
		
		//Simple validation to make sure user entered something
		//Add your own error checking here with JS, but also do some error checking with PHP.
		//If error found, add hightlight class to the text field
		if (name.val()=='') {
			name.addClass('error');
			returnError = true;
		} else name.removeClass('error');
		
		if (email.val()=='') {
			email.addClass('error');
			returnError = true;
		} else email.removeClass('error');		
		
		if(!regx.test(email.val())){
          email.addClass('error');
          returnError = true;
		} else email.removeClass('error');

		// Highlight all error fields, then quit.
		if(returnError == true){
			return false;	
		}
		
		//organize the data
		
		var data = 'name=' + name.val() + '&email=' + email.val();

		//disabled all the text fields
		$('.text').attr('disabled','true');
		
		//show the loading sign
		$('.loading').show();
		
		//start the ajax
		$.ajax({
			//this is the php file that processes the data and sends email
			url: "/g/",
			
			//GET method is used
			type: "POST",

			//pass the data			
			data: data,		
			
			//Do not cache the page
			cache: false,
		});
		
		//cancel the submit button default behaviours
		return false;
	});	
});	