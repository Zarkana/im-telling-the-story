import React, { Component } from "react";
import {
    HashRouter
  } from "react-router-dom";
import Header from "./Header";
import Content from "./Content/Content";
import Footer from "./Footer";
import ProgressBar from "./ProgressBar/ProgressBar";
import PrimaryNav from "./PrimaryNav/PrimaryNav";

class Main extends Component {
  render() {
    return (        
        <HashRouter>
            <div className="container">
                <Header/>
                <ProgressBar/>
                <Content/>
                <PrimaryNav/>
                <Footer/>
            </div>
        </HashRouter>
    );
  }
}
 
export default Main;