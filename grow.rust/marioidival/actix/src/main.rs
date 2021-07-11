pub mod core {
    pub mod domain {
        pub struct Country {
            pub name: String,
            pub indicator: String,
            pub value: f64,
            pub year: u64,
        }
    }
    pub mod ports {
        trait Country {}

        trait CountryRepository {
            fn growth_info(&self) -> crate::core::domain::Country;
            fn size(&self) -> usize;
            fn update_growth(&self);
            fn create_country_growth_info(&self);
        }
    }

    pub mod services {
        /// get status of processing (LoadGrowthInformation)
        pub struct StatusProcessService {}
        /// get some growth unit information of growth database
        pub struct GrowthInformationService {}
        /// save all growth information from file
        pub struct LoadGrowthInformationService {}
        /// update some growth unit information
        pub struct UpdateGrowthInformationService {}
        /// remove country growth information
        pub struct RemoveGrowthInformationService {}
    }
}

pub mod adapters {
    pub mod memory {}
    pub mod http {
        pub mod controllers {}

        pub struct Server;
    }
}

pub mod config {
}

fn main() {
    println!("Hello, world!");
}
