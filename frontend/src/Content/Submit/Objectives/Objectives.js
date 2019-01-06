import React, { Component } from "react";
 
class Objectives extends Component {
  render() {
    return (
      <div className="objectives content-holder shift-text-down-children">
        <p className="objective"><span className="points">+10</span> If the word <span className="not-met">banana</span> is used at least once</p>
        <p className="objective"><span className="points">+10</span> If less than <span className="met">4</span> sentences are used</p>
        <p className="objective"><span className="points">+10</span> If <span className="not-met">2</span> semi colons are used</p>
        <p className="total-points"><span className="points">=30</span> Points if it wins</p>
      </div>
    );
  }
}
 
export default Objectives;