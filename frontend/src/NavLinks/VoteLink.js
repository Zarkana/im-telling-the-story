import React, { Component } from "react";
import {
  NavLink,
} from "react-router-dom"; 

class VoteLink extends Component {
  
  render() {
    return (      
        <NavLink exact to="/"> &larr; Home</NavLink>
        // <li><NavLink to="/game">Game</NavLink></li>
        // <NavLink to="/vote">Vote --></NavLink>
    );
  }
}

export default VoteLink;
