pub mod pipe;
pub mod bird;


use std::io;
use crate::screen::Screen;
use self::{bird::Bird, pipe::Pipe};

pub trait RenderAble{
    fn render(&self, screen: &dyn Screen);
}

pub struct Game <'a>{
    pub screen: &'a dyn Screen,
    pub bird: Bird,
    pub pipe: Pipe,
}

impl<'a> Game<'a>{
    
    pub fn new(screen: &'a dyn Screen) -> Self {
        return Self{
            screen,
            bird: Bird::new(),
            pipe: Pipe::new()
        }
    }

    pub fn start(&mut self) -> io::Result<()>{
        self.bird.render(self.screen);
        self.pipe.render(self.screen);
        return Ok(());
    }
}

