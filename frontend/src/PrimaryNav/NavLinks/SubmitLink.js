import React, { Component } from "react";
import {
  NavLink,
} from "react-router-dom"; 

class SubmitLink extends Component {
  
  render() {
    return (
      <div className="shift-text-down">            
        <NavLink to="/submit">Add Submission &rarr;</NavLink>
      </div>
    );
  }
}

export default SubmitLink;
