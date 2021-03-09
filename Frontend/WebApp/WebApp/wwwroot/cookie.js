function setCookie(cname, cvalue, exdays, exhours, exminutes) {
    var d = new Date();
    if (exdays == 0) {
        if (exhours == 0) {
            d.setTime(d.getTime() + (exminutes * 60 * 1000)); //Minutes to ms
        } else {
            d.setTime(d.getTime() + (exhours * 60 * 60 * 1000)); //Hours to ms
        }
    } else {
        d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000)); //Days to ms
    }
    var expires = "expires=" + d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/" + ";SameSite=Lax";
}
function getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}