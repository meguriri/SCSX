$(document).ready(function () {
    $.ajax({
        type: "get",
        url: "/order/myOrderList",
        dataType: 'json',
        success: function (res) {
            if (res.msg === "ok") {
                let obj=eval(res.list)
                let str=''
                console.log("obj:",obj)
                for(let i=0;i<obj.length;i++){
                    str+='<tr>\n' +
                        '                <td><font color="#ff4e00">'+obj[i].id+'</font></td>\n' +
                        '                <td>'+obj[i].time+'</td>\n' +
                        '                <td>'+obj[i].address+'</td>\n' +
                        '              </tr>'
                }
                $('#title').after(str)
            }
        },
    })
})