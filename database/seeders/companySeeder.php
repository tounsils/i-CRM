<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use App\Models\company;  // to be added

class companySeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
      // \App\Models\User::factory(10)->create();
      company::factory()
      ->count(13)
      ->create();
    }
}
