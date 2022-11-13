# Ascii-Art-Web
___

### Table of Contents

---

- [Description](#description)
- [Authors](#authors)
- [How-To-Use](#usage)
- [Implementation details](#algorithm)

---

## Description
    This project consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of ascii-art. 
    Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. What we mean by a graphic representation using ASCII, is to write the string received using ASCII characters like that:
    
![sample text](https://imageup.ru/img51/3953409/screenshot-2022-06-11-at-15-47-39-ascii-art-web.png)


---

## Authors
Markhabat (@Markhabat) & Nurzhas (@Nurzhas)

---

## Usage
1. To run the server: go run . OR go run main.go
2. Once the server is started, open the page in the browser using the correct port or you can use the link that will be in the terminal.
3. Enter any text in the input area on the page.
4. Choose the banner and click the button to get the ascii-art.

---

## Algorithm
    Implemented the following HTTP endpoints:

    1. GET /: Sends HTML response, the main page. 
    The function responsible for this in our program is Handler(w http.ResponseWriter, r *http.Request) in internal/delivery/mainPageHandler.go

    2. POST /ascii-art: that sends data to Go server (text and a banner)
    The function responsible for this in our program is SendArt(w http.ResponseWriter, r *http.Request) in internal/delivery/router.go
    This function takes the banner and text from the browser and processes this text using the function Getresult(banner, text string) (string, error) to get the Ascii-Art.
