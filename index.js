var url = "http://localhost:8889/";

async function GetMessage() {
  const j = await fetch(url, {
    method: "GET"
  })
  fetch(url)
    .then(response => response.json())
    .then(j => document.getElementById("message").innerHTML = j.message)
}

function PostMessage(m) {
  const param = {
    method: "POST",
    headers: {
      "Content-Type": "application/json; charset=utf-8"
    },
    // リクエストボディ
    body: JSON.stringify({
      message : m
    })
  };
  
  fetch(url, param)
    .then(function (response) {
      console.log(response.text)
    })
    .catch(err => {
      console.error(err)
    })
}

