package main

import (
	"fmt"
	"net/http"
)

func pageMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<!DOCTYPE html5>
		<html>
		<head>
		  <meta name="viewport" content="width=device-width, initial-scale=1.0">
		  <title>VueLamp (beta)</title>
		  <link rel="stylesheet" href="style_sito(1).css">
		</head>
		<body>
		  <div class="hero">
		    <nav>
		      <img src="images/menu.png" class="">
		      <img src="images/logo.png" class="logo-img">
		      <ul>
		        <li><a href="">Latest</a></li>
		        <li><a href="">Modern</a></li>
		        <li><a href="">Contemporary</a></li>
		        <li><a href="">Affordable</a></li>
		      </ul>
		      <button type="button" onclick="toggleBtn()" id="btn" name="button1"><span></span></button>
		      <i class="arrow"></i>
		    </nav>
		      <div class="lamp-container">
		        <img src="images/lamp.png" class="lamp-img">
		        <img src="images/light.png" class="light-img" id="light">
		      </div>
		      <div class="text-container">
		        <h1>Lighting</h1>
		        <p><br>The first lamp from our company, from home use to office use.</p>
		        <a href="">Check the collection</a>
		        <div class="control">
		          <p>04</p>
		          <div class="line"><span></span></div>
		          <p>05</p>
		        </div>
		      </div>
		  </div>
		  <script>

		    var btn = document.getElementById("btn");
		    var light = document.getElementById("light");

		    function toggleBtn() {
		      btn.classList.toggle("active");
		      light.classList.toggle("on");
		    }
		  </script>
		</body>
		</html>
	`))
}

func main() {
	http.HandleFunc("/", pageMain)
	fmt.Println("Listening on http://localhost:8000/")
	http.ListenAndServe(":8000", nil)
}