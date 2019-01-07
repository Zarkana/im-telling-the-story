import React, { Component } from "react";

class Footer extends Component {

  constructor(props) {
    super(props);
      this.state = {
        username: "Brian",
        loggedIn: false,//TODO: Move to higher component
      };
  }

  getFooterLogin() {
    if(this.state.loggedIn){
      return this.state.username;
    }else{
      return "SIGN IN";
    }
  }

  render() {
    return (
      <footer>
          <div className="left shift-text-down">{this.getFooterLogin()}</div>
          <div className="center shift-text-down">
          </div>
          <div className="right shift-text-down">GITHUB</div>
      </footer>
    );
  }
}

export default Footer;

