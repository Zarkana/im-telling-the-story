import React, { Component } from "react";
import Requirements from "./Requirements/Requirements";
import SubmissionField from "./SubmissionField/SubmissionField";
import Story from "../../Story/Story"; 
import './Submit.css';

class Submit extends Component {
  render() {
    return (
      <div className="submit">
        <Requirements/>
        <Story/>
        <SubmissionField/>
      </div>
    );
  }
}
 
export default Submit;