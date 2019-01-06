import React, { Component } from "react";
 
class Footer extends Component {
  render() {
    return (
      <footer>
          <div className="left shift-text-down">SIGN IN</div>
          <div className="center shift-text-down">
            <button>
              Jimmy
            </button>
            <button>
              Rob
            </button>
            <button>
              Sally
            </button>
            <button>
              Martha
            </button>
            <button>
              Logout
            </button>
          </div>
          <div className="right shift-text-down">GITHUB</div>
      </footer>
    );
  }
}

export default Footer;

