import React, { Component } from "react";
import './submission.css';

class Submission extends Component {
  render() {
    const submission = this.props.submission;
    // const text = submission.text;

    return (
      <span>{submission.text}</span>
    );
  }}
export default Submission;