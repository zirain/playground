import React, { Component } from 'react';
import 'antd/dist/antd.css';
import { Carousel } from 'antd';

export class Home extends Component {
    render() {
        var settings = {
            className: "center",
            centerMode: true,
            infinite: true,
            centerPadding: "60px",
            slidesToShow: 3,
            speed: 500,
            autoplay: true
        };
        return (
            <div>
                <h1>Carousel Demo</h1>
                <Carousel {...settings}>
                    <div>
                        <h3>1</h3>
                    </div>
                    <div>
                        <h3>2</h3>
                    </div>
                    <div>
                        <h3>3</h3>
                    </div>
                    <div>
                        <h3>4</h3>
                    </div>
                </Carousel>
            </div>
        );
    }
}
