import React, { Component } from "react";
import Objective from './Objective/Objective';
import './objective.css';

class ObjectiveList extends Component {
  render() {
    const rows = [];    
    dummyObjectives.forEach((objective) => {
      rows.push(
        <Objective
          objective={objective}
          key={objective.id} />
      );
    });
    return (
      <div className="objectives content-holder shift-text-down-children">
        {rows}
        <p className="total-points"><span className="points">=30</span> Points if wins voting</p>
      </div>
    );
  }}
export default ObjectiveList;

var dummyObjectives = [
  {         
    id: '1',
    points: '10',
    pre: "If the word ",
    qualification: "banana",
    post: " is used at least once"    
  },
  {         
    id: '2',
    points: '10',
    pre: " If less than ",
    qualification: "4",
    post: " sentences are used"
  },
  {       
    id: '3',  
    points: '10',
    pre: "If at least ",
    qualification: "2",
    post: " semi-colons are used"    
  },    
];