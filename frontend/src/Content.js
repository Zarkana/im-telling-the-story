import React, { Component } from "react";
import { Route } from "react-router-dom";
import Home from "./Home";
import Game from "./Game";
import Vote from "./Vote";

class Content extends Component {
  render() {
    return (                
        <div className="content">
            <Route exact path="/" component={Home}/>
            <Route path="/game" component={Game}/>
            <Route path="/vote" component={Vote}/>
        </div>
    );
  }
}

export default Content;
