import React, { Component } from "react";
import { Route } from "react-router-dom";
import Home from "./Home/Home";
import Game from "../Game/Game";
import Vote from "./Vote";
import Submit from "./Submit";
import './content.css';

class Content extends Component {
  render() {
    return (                
        <div className="content">
            <Route exact path="/" component={Home}/>
            <Route path="/game" component={Game}/>
            <Route path="/vote" component={Vote}/>
            <Route path="/submit" component={Submit}/>
        </div>
    );
  }
}

export default Content;
