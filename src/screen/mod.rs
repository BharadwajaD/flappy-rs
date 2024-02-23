use std::io::{self, Stdout};

use ratatui::{
    backend::CrosstermBackend,
    layout::Rect,
    widgets::{Block, Borders},
    Terminal,
};

use crate::game::shape::SRectangle;

pub struct CrossTermDisplay {
    pub terminal: Terminal<CrosstermBackend<Stdout>>,
}

impl CrossTermDisplay {
    pub fn new(terminal: Terminal<CrosstermBackend<Stdout>>) -> Self {
        return Self { terminal };
    }
}

impl Screen for CrossTermDisplay {
    fn get_rectangle(&self, rect: &SRectangle) -> Rect {
        return Rect::new(rect.x, rect.y, rect.width, rect.height);
    }

    fn draw(&mut self, area: Rect) -> io::Result<()> {
        self.terminal.draw(|frame| {
            frame.render_widget(Block::default().borders(Borders::ALL), area);
        })?;
        return Ok(());
    }

    fn get_fsize(&self) -> (u16, u16) {
        let fsize = self.terminal.size().unwrap();
        return (fsize.width, fsize.height);
    }
}

pub trait Screen {
    fn get_fsize(&self) -> (u16, u16);
    fn get_rectangle(&self, rect: &SRectangle) -> Rect;
    /// calls ui.draw
    fn draw(&mut self, area: Rect) -> io::Result<()>;
}
