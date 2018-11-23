import React, { Component } from "react";
import { Route } from "react-router-dom";
import VoteLink  from "./NavLinks/VoteLink";
import SubmitLink  from "./NavLinks/SubmitLink";
import Home from "./Home";
import {
  NavLink,
} from "react-router-dom"; 

class PrimaryNav extends Component {
  
  render() {
    return (
      <div className="primary-nav">
      {/* <ul className="header">
          <li><NavLink exact to="/">Home</NavLink></li>
          <li><NavLink to="/game">Game</NavLink></li>
          <li><NavLink to="/vote">Vote --></NavLink></li>
      </ul> */}
      <Route exact path="/" component={SubmitLink}/>
      <Route path="/vote" component={VoteLink}/>
      
  </div>
    );
  }
}

export default PrimaryNav;
