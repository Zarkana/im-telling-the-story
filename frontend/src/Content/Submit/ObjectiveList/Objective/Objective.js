import React, { Component } from "react";

class Objective extends Component {
  render() {
    const objective = this.props.objective;

    return (
      <p className="objective" key={objective.id}>
          <span className="points">+{objective.points}</span> {objective.pre} <span className="not-met">{objective.qualification}</span> {objective.post}
      </p>
    );
  }}
export default Objective;