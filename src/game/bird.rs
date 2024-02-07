use crate::screen::Screen;

use super::RenderAble;

pub struct Bird{
}

impl Bird {
    pub fn new() -> Self {
        return Self {  };
    }
}

impl RenderAble for Bird {
    fn render(&self, screen: &dyn Screen) {
        screen.get_rectangle();
    }
}
