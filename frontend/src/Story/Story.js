import React, { Component } from "react";
import Submission from "./Submission/Submission";
import './story.css';

class Story extends Component {
  render() {
    const submissions = [];
    dummySubmissions.forEach((submission) => {
      submissions.push(
        <Submission
          submission={submission}
          key={submission.id} />
      );
    });
    return (
      <div className="story content-holder">        
        {submissions}
      </div>
    );
  }}
export default Story;

var dummySubmissions = [
  {         
    id: '1',
    text: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. ',    
  },
  {         
    id: '2',
    text: 'Dolorem laboriosam quae incidunt, error minus quos molestias suscipit in corporis repudiandae mollitia. ',    
  },
  {       
    id: '3',  
    text: ' Non explicabo possimus dolore praesentium reprehenderit accusantium iste ea vero labore unde veniam beatae consectetur qui, voluptatem repellendus asperiores quas repellat quibusdam?',    
  },    
];