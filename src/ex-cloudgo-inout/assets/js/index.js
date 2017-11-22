(function() {
    var clock;
    var waitTime = 5;
    
    $(function() {
        $("input[type='text']").blur(checkValidity);
        $("[name='signin']").submit(submitForm);
        $("[type='reset']").click(clearImages);
        if(window.location.href != "http://localhost:8080/register") window.location.href = "http://localhost:8080/register";
        if($("input[type='text']").val() != "") $("input[type='text']").blur();
    });
    
    function checkValidity() {    //检查合法性
        var value = $(this).val();
        if (value == "") return;
        var name = $(this).attr("name");
        switch(name) {
            case "username": 
                /^[a-zA-Z]\w{5,17}$/.test(value) ? showValidity(true, name) : showValidity(false, name);
                break;
            case "studentID":
                /^[1-9]\d{7}$/.test(value) ? showValidity(true, name) : showValidity(false, name);
                break;
            case "phone":
                /^[1-9]\d{10}$/.test(value) ? showValidity(true, name) : showValidity(false, name);
                break;
            case "mail":
                /^[a-zA-Z_\-]+@(([a-zA-Z_\-])+\.)+[a-zA-Z]{2,4}$/.test(value) ? showValidity(true, name) : showValidity(false, name);
                break;
        }
    }
    
    function showValidity(flags, name) {    //显示是否合法的图标
        var icon = $("img." + name);
        icon.attr("src", flags ? "./images/correct.png" : "./images/error.png");
        icon.css("opacity", 1);
    }
    
    function clearImages() {
        $("img").each(function() {
            $(this).css("opacity", 0);
        })
    }
    
    function isAllCorrect() {    //全部合法才能提交
        var corrcetPath = "./images/correct.png";
        var flag = true;
        $("img").each(function() {
            if($(this).attr("src") != corrcetPath) flag = false;
        });
        return flag;
    }
    
    function submitForm() {
        if(!isAllCorrect()) {    //不合要求则阻止提交
            $("#result").text("存在非法格式！");
            return false;
        }
           
        return ture; //submit提交
    } 
})();
