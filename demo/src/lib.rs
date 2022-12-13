
pub mod demo {
    use std::panic;
    /// # Safety
    #[no_mangle]
    pub unsafe extern "C" fn trying() {
        panic::catch_unwind(|| panic!("boom!"));
    }
}
