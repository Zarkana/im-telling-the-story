import React, { Component } from "react";
import Story from "../Story/Story";
import SubmissionCount from "../SubmissionCount/SubmissionCount";

class Home extends Component {
  render() {
    return (
      <div className="home-content">
        <Story></Story>
        <SubmissionCount></SubmissionCount>
      </div>
    );
  }
}
 
export default Home;