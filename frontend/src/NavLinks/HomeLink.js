import React, { Component } from "react";
import {
  NavLink,
} from "react-router-dom"; 

class HomeLink extends Component {
  
  render() {
    return (      
        <NavLink exact to="/">&larr; Home</NavLink>
    );
  }
}

export default HomeLink;
