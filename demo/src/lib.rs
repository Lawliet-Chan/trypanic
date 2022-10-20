
pub mod demo {
    /// # Safety
    #[no_mangle]
    pub unsafe extern "C" fn trying() {
        panic!("boom!")
    }
}
