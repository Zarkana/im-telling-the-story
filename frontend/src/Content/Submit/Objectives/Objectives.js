import React, { Component } from "react";

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

class Objectives extends Component {

  constructor(){
    super();
    this.state = {
      objectives: [],
    };
  }

  componentDidMount(){
    const objectivess = dummyObjectives.map((objective) => 
    <p className="objective" key={objective.id}>
      <span className="points">+{objective.points}</span> {objective.pre} <span className="not-met">{objective.qualification}</span> {objective.post}
    </p>
    );
    this.setState({objectives: objectivess});
    console.log("state", this.state.objectives);
  }

  render() {
        const data =[{"name":"test1"},{"name":"test2"}];
      const listItems = data.map((d) => <li key={d.name}>{d.name}</li>);

    return (
      <div className="objectives content-holder shift-text-down-children">
      {this.state.objectives }
        <p className="total-points"><span className="points">=30</span> Points if wins voting</p>
      </div>
    );
  }

}
 
export default Objectives;