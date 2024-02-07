use std::io::Stdout;

use ratatui::{Terminal, backend:: CrosstermBackend};


pub struct CrossTermDisplay{
    pub terminal: Terminal<CrosstermBackend<Stdout>>
}

impl CrossTermDisplay {

    pub fn new(terminal: Terminal<CrosstermBackend<Stdout>>) -> Self{
        return Self{
            terminal
        }
    }

}

impl Screen for CrossTermDisplay {
    fn get_rectangle(&self) {
        println!("rectangle here");
    }

    fn draw(&mut self) {
        self.terminal.draw(|_|{});
    }
}

pub trait Screen {
    fn get_rectangle(&self);
    fn draw(&mut self);
}
