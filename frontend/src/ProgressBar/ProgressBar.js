import React, { Component } from "react";
import './progress-bar.css';

class ProgressBar extends Component {
  render() {
    return (
      <div className="progress-bar">
          <div className="left side-content"><span></span>31%<span className="less-emphasis"> - Till Complete</span></div>
          <div className="center round-number">Round 1</div>
          <div className="right side-content"><span className="less-emphasis">Voting - </span> 12:43:26</div>
      </div>
    );
  }
}

export default ProgressBar;
