{{template "../public/header.tpl"}}
<script type="text/javascript" src="/static/js/jquery.Huploadify.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/Huploadify.css">



<script type="text/javascript">
$(function(){
	$('#upload').Huploadify({
		auto:true,
		fileTypeExts:'*.jpg;*.png;*.exe;*.jpeg;*.JPG',
		multi:true,
		formData:{key:123456,key2:'vvvv'},
		fileSizeLimit:9999,
		showUploadedPercent:true,//是否实时显示上传的百分比，如20%
		showUploadedSize:true,
		removeTimeout:9999999,
		uploader:'/picmanager/init/preview',
		onUploadStart:function(){
			//alert('开始上传');
			},
		onInit:function(){
			//alert('初始化');
			},
		onUploadComplete:function(file){
			//console.log(file.name)
			//$("#showpic").append("<img src='' width='40px' height='40px' />")
			//alert('上传完成');
			},
		onDelete:function(file){
			console.log('删除的文件：'+file);
			console.log(file);
		},
		onUploadSuccess: function(file,data){
			var obj = JSON.parse(data)
			
			if (obj.total > 0) {
				for(i= 0;i<obj.total;i++) {
					var picspan = ""
					console.log(obj.rows[i].pathinfo)
					picspan = "<span><img src='/" + obj.rows[i].pathinfo + "' width='40px' height='40px' /><input name='pic[]' type='hidden' value='" + obj.rows[i].pathinfo + "' /></span>"
					$("#showpic").append(picspan)
					
				}
			}
			
/*                    alert( 'id: ' + file.id
　                          + ' - 索引: ' + file.index
　　　　　　　　　　　　　　　　+ ' - 文件名: ' + file.name
　　　　　　　　　　　　　　　　+ ' - 文件大小: ' + file.size
　　　　　　　　　　　　　　　　+ ' - 类型: ' + file.type
　　　　　　　　　　　　　　　　+ ' - 创建日期: ' + file.creationdate
　　　　　　　　　　　　　　　　+ ' - 修改日期: ' + file.modificationdate
　　　　　　　　　　　　　　　　+ ' - 文件状态: ' + file.filestatus
　　　　　　　　　　　　　　　　+ ' - 服务器端消息: ' + data);
*/
                }
		});
	});
</script>
</head>

<body>
<div id="upload"></div>


<form action="/picmanager/init/save">

<div id="showpic"></div>

<input type="submit" value="上传" />
</form>

