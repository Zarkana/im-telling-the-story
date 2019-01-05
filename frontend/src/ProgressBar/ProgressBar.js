import React, { Component } from "react";
import './progress-bar.css';

class ProgressBar extends Component {
  render() {
    return (
      <div className="progress-bar">
          <div className="left shift-text-down">31%</div>
          <div className="center shift-text-down">The Story Thus Far</div>
          <div className="right shift-text-down">12:43:26</div>
      </div>
    );
  }
}

export default ProgressBar;
