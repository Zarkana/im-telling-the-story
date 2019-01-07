import React, { Component } from "react";
import ObjectiveList from "./ObjectiveList/ObjectiveList";
import SubmissionField from "./SubmissionField/SubmissionField";
import Story from "../../Story/Story"; 
import './Submit.css';

class Submit extends Component {
  render() {
    return (
      <div className="submit">
        <ObjectiveList/>
        <Story/>
        <SubmissionField/>
      </div>
    );
  }
}
 
export default Submit;