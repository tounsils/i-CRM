<?php

namespace Database\Factories;

use App\Models\company;
use Illuminate\Database\Eloquent\Factories\Factory;

class companyFactory extends Factory
{
    /**
     * The name of the factory's corresponding model.
     *
     * @var string
     */
    protected $model = company::class;

    /**
     * Define the model's default state.
     *
     * @return array
     */
    public function definition()
    {
        return [
            'company_name' => $this->faker->company(),
            'email' => $this->faker->unique()->safeEmail(),
            'phone' => $this->faker->phoneNumber,
        ];
    }
}

