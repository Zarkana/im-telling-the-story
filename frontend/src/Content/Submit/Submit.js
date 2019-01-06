import React, { Component } from "react";
import Objectives from "./Objectives/Objectives";
import SubmissionField from "./SubmissionField/SubmissionField";
import Story from "../../Story/Story"; 
import './Submit.css';

class Submit extends Component {
  render() {
    return (
      <div className="submit">
        <Objectives/>
        <Story/>
        <SubmissionField/>
      </div>
    );
  }
}
 
export default Submit;