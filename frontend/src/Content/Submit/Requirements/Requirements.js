import React, { Component } from "react";
 
class Requirements extends Component {
  render() {
    return (
      <div className="requirements content-holder shift-text-down-children">
        <p className="requirement"><span className="points">+10</span> Must use the word <span className="not-met">banana</span> at least once</p>
        <p className="requirement"><span className="points">+10</span> Must be less than <span className="met">4</span> sentences</p>
        <p className="requirement"><span className="points">+10</span> Must use <span className="not-met">2</span> semi colons</p>
        <p className="total-points"><span className="points">= 30</span> POINTS</p>
      </div>
    );
  }
}
 
export default Requirements;