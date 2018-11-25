import React, { Component } from "react";
import {
    HashRouter
  } from "react-router-dom";
import Header from "./Header";
import Content from "./Content";
import Footer from "./Footer";
import ProgressBar from "./ProgressBar";
import PrimaryNav from "./PrimaryNav";

class Main extends Component {
  render() {
    return (        
        <HashRouter>
            <div className="container">
                <Header></Header>
                <ProgressBar></ProgressBar>
                <Content></Content>
                <PrimaryNav></PrimaryNav>
                <Footer></Footer>
            </div>
        </HashRouter>
    );
  }
}
 
export default Main;