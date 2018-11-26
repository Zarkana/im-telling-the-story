import React, { Component } from "react";
import {
  NavLink,
} from "react-router-dom"; 

class SubmitLink extends Component {
  
  render() {
    return (
        <div>            
          <NavLink to="/submit">Submit &rarr;</NavLink>
      </div>
    );
  }
}

export default SubmitLink;
