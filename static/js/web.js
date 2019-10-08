function getUserNameFromCookie() {
    if (document.cookie.length > 0) {
        let index1 = document.cookie.indexOf("username");
        let index2 = document.cookie.indexOf("$");
        if (index1 >= 0 && index2 >= 0) {
            return document.cookie.substring(index1, index2).split("=")[1];
        }
    }
}