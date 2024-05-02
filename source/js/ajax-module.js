// global page ready
$(document).ready(function(){
	$(".tab-content a").each(function(){
		if(this.href == MODULE_URL){
			$(this).parent().addClass('active');
			$(this).parent().parent().parent().addClass('in active');
			content = $(this).parent().parent().parent().attr('id');
			$(".nav-tabs a").each(function(){
				if(this.href == MODULE_URL+'#'+content){
					$(this).parent().addClass('active');
				}
			});
		};
	});
	// generateBreadcrumb();
});

// global confirm box
// $(document).on('click', '.confirm-action', function(e){
    
//     var thisObj = $(this);
    
//     if(! thisObj.data('statestatus')){
//         e.preventDefault();
		
// 		showConfirm().then((result) => {
// 			if(result.value){
// 				// set status pada data
// 				thisObj.attr('data-statestatus', 'OK');
// 				thisObj[0].click();
// 			}
// 		});
//     }
// });

// generate page breadcrumb
// function generateBreadcrumb(url_data = null){
// 	$.get(MODULE_URL + "/ajax_page_breadcrumb", {url: url_data}).done(function(data){
// 		$('.breadcrumb').html(data);
// 	});
// }

// ajax module pages
function showModulePage(f_url, f_method, f_setting, f_data, f_page_url){
	mod_content = null;
	mod_spinner = null;
	mod_processData = true;
	mod_contentType = "application/x-www-form-urlencoded; charset=UTF-8";

	if(f_setting['content']){
		mod_content = f_setting['content'];
	}if(f_setting['spinner']){
		mod_spinner = f_setting['spinner'];
	}if(f_setting['post']){
		mod_processData = !f_setting['post'];
		mod_contentType = false;
	}

	// update breadcrumb
	// generateBreadcrumb(f_page_url);

	return $.ajax({
		url      : f_url,
		cache: true,
        contentType: mod_contentType,
        processData: mod_processData,
		type     : f_method,
		data     : f_data,
		dataType : "JSON",
		beforeSend : function(){
			moduleStatusInfo(f_setting['status'], mod_content, mod_spinner, "process");
		},
		error: function(xhr){
			moduleStatusInfo(f_setting['status'], mod_content, mod_spinner, "error", xhr.statusText);
		},
		success: function(data){
			if(data.status === true){
				if(f_page_url != null){
					moduleSetURL(f_page_url, data.title + " - " + APP_NAME);
				}

				moduleStatusInfo(f_setting['status'], mod_content, mod_spinner, "ok");
			}else{
				moduleStatusInfo(f_setting['status'], mod_content, mod_spinner, "error");
			}
		},
		complete: function(){
			$('.data-table').DataTable();
		}
	});
}

function showModulePage2(f_url, f_method, f_setting, f_data, f_page_url){
	mod_content = null;
	mod_spinner = null;

	if(f_setting['content']){
		mod_content = f_setting['content'];
	}if(f_setting['spinner']){
		mod_spinner = f_setting['spinner'];
	}

	// update breadcrumb
	generateBreadcrumb(f_page_url);

	return $.ajax({
		url      : f_url,
		type     : f_method,
		data     : f_data,
		dataType : "JSON",
		processData:false,
        contentType:false,
        cache:false,
        async:true,
		beforeSend : function(){
			moduleStatusInfo(f_setting['status'], mod_content, mod_spinner, "process");
		},
		error: function(xhr){
			moduleStatusInfo(f_setting['status'], mod_content, mod_spinner, "error", xhr.statusText);
		},
		success: function(data){
			if(data.status === true){
				if(f_page_url != null){
					moduleSetURL(f_page_url, data.title + " - " + APP_NAME);
				}

				moduleStatusInfo(f_setting['status'], mod_content, mod_spinner, "ok");
			}else{
				moduleStatusInfo(f_setting['status'], mod_content, mod_spinner, "error");
			}
		},
		complete: function(){
			$('.selectpicker').chosen();
			$('.datatable').DataTable();

		}
	});
}

// sweetalert
function showAlert(is_success = true, msg = null){
	if(is_success == true){
		alTitle	= "Aksi Berhasil";
		alType 	= "success";
	}else{
		alTitle = "Aksi Gagal";
		alType 	= "error";
	}
	
	return swal({
        title: alTitle,
        html: msg,
        type: alType,
        confirmButtonClass: 'btn btn-primary',
        timer: 2500
    });
}

// sweetalert confirmation
function showConfirm(){
	return swal({
		title: "Konfirmasi Aksi",
		html: "Apakah Anda Yakin dengan Aksi Ini?",
		type: "warning",
		showCancelButton: true,
		confirmButtonColor: '#3085d6',
		cancelButtonColor: '#d33',
		confirmButtonText: 'Ya',
		confirmButtonClass: 'btn btn-primary',
		cancelButtonClass: 'btn btn-danger',
		cancelButtonText: 'Tidak'
	});
}

// module status
function moduleStatusInfo(element, element_main, element_spinner, type, message){
	// cek apakah main element kosong atau tidak
	if(element_main !== null){
		element_main.empty();
	}

	switch(type){
		case "process":
			if(element_spinner !== null){
				element_spinner.html('<center class="text-muted mt-3"><i class="fa fa-spinner fa-spin font-24 text-success"></i></center>');
			}

			element.removeClass("text-danger");                
			element.html("Loading...");
			break;
		case "ok":
			if(element_spinner !== null){
				element_spinner.empty();
			}

			element.removeClass("text-danger");
			element.html("Ready");
			break;
		case "error":
			if(element_spinner !== null){
				element_spinner.empty();
			}

			element.addClass("text-danger");
			element.html('Error: ' + message);
			break;
	}
}

// set module URL & title
function moduleSetURL(url, title){
	window.history.pushState(null, title, url);
	
	if(title){
		document.title = title;
	}
}