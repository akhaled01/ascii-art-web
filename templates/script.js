// This function will make the request to the back-end and will modify the html
function loadDoc() {
    const xhttp = new XMLHttpRequest();
    let text = document.querySelector(".input")
    let banner = document.querySelector('input[name="Banner"]:checked');
    let Color = document.querySelector(".newColor")
    xhttp.onload = function () {
        var parsedData = JSON.parse(this.responseText);                 // Parse JSON
        //var formattedText = parsedData.Result.replace(/\n/g, "<br>");   // Replace the lines with <br>
        document.getElementById("art").innerHTML = parsedData.Result;
        document.getElementById("art").style.color = parsedData.ApplyColor;
    }
    xhttp.open("POST", "ascii-art");
    xhttp.setRequestHeader("Content-Type", "application/json; charset=utf-8");
    
    const body = {
        Text: text.value,
        Banner: banner.value,
        Newcolor: Color.value
    };

    xhttp.send(JSON.stringify(body));
}