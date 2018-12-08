import React, { Component } from "react";
import {
  NavLink,
} from "react-router-dom"; 

class HomeLink extends Component {
  
  render() {
    return ( 
      <div className="shift-text-down">
        <NavLink exact to="/">&larr; The Story Thus Far</NavLink>
      </div>
    );
  }
}

export default HomeLink;
