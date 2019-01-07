import React, { Component } from "react";
import './progress-bar.css';

class ProgressBar extends Component {
  constructor(props) {
    super(props);
    this.state = {
      round: "1",
      percentage: "31%",
      stage: "Voting",
      timer: "12:43:23",
    };
  }
  render() {
    return (
      <div className="progress-bar">
        <div className="left side-content"><span></span>{this.state.percentage}<span className="less-emphasis"> - Till Complete</span></div>
        <div className="center round-number">Round {this.state.round}</div>
        <div className="right side-content"><span className="less-emphasis">{this.state.stage} -</span> {this.state.timer}</div>
      </div>
    );
  }
}
export default ProgressBar;
