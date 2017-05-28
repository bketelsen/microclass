
$(() => {

});
// src/index.js
var m = require("mithril")

var UserList = require("./views/UserList")
var Login= require("./views/Login")
// The profile model
var auth = require('./models/Auth')

if (auth.authenticate()) {
  AppRouter()
} else {
  LoginRouter()
};

function AppRouter() {
  document.body.id = 'authenticated'
  m.route(document.body, '/', {
   '/': UserList 
  })
}

function LoginRouter() {
  document.body.id = 'login'
  m.route(document.body, '/', {
    '/': Login
  })
}